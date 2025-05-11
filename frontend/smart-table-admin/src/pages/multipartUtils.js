export function indexOf(buf, sub, from = 0) {
  for (let i = from; i <= buf.length - sub.length; i++) {
    let ok = true;
    for (let j = 0; j < sub.length; j++) {
      if (buf[i + j] !== sub[j]) { ok = false; break; }
    }
    if (ok) return i;
  }
  return -1;
}

export function parseMixed(bodyBuf, boundary) {
  const enc = new TextEncoder();
  const bnd = enc.encode(`--${boundary}`);
  let pos = indexOf(bodyBuf, bnd);
  if (pos < 0) return [];
  pos += bnd.length;
  const parts = [];

  while (true) {
    if (bodyBuf[pos] === 45 && bodyBuf[pos + 1] === 45) break; 
    if (bodyBuf[pos] === 13 && bodyBuf[pos + 1] === 10) pos += 2; 

    const next = indexOf(bodyBuf, bnd, pos);
    if (next < 0) break;

    let chunk = bodyBuf.subarray(pos, next);
    if (chunk[chunk.length - 2] === 13 && chunk[chunk.length - 1] === 10) {
      chunk = chunk.subarray(0, chunk.length - 2);
    }

    const sep = indexOf(chunk, enc.encode('\r\n\r\n'));
    const headBuf = chunk.subarray(0, sep);
    const dataBuf = chunk.subarray(sep + 4);

    const headText = new TextDecoder().decode(headBuf);
    const headers = {};
    headText
      .split('\r\n')
      .filter(line => line.includes(':'))
      .forEach(line => {
        const [k, ...rest] = line.split(':');
        headers[k.trim().toLowerCase()] = rest.join(':').trim();
      });

    const cd = headers['content-disposition'] || '';
    const nameMatch = cd.match(/name="([^"]+)"/i);
    const fileMatch = cd.match(/filename="([^"]+)"/i);

    parts.push({
      name: nameMatch?.[1] || null,
      filename: fileMatch?.[1] || null,
      type: headers['content-type'] || null,
      data: dataBuf
    });

    pos = next + bnd.length;
  }

  return parts;
}

export async function handleMultipartResponse(response, listField = 'dish_list') {
    const ct = response.headers.get("Content-Type") || "";
    const [, boundary] = ct.match(/boundary="?([^";]+)"?/) || [];
    if (!boundary) throw new Error("Не удалось вытащить boundary");

    const buf = new Uint8Array(await response.arrayBuffer());
    const parts = parseMixed(buf, boundary);

    const jsonPart = parts.find(p => p.type === "application/json");
    if (!jsonPart) throw new Error("JSON часть не найдена");
    
    const json = JSON.parse(new TextDecoder().decode(jsonPart.data));
    const list = Array.isArray(json) ? json : json[listField] || [];

    const imagesMap = {};
    parts.filter(p => p.filename).forEach(p => {
        const blob = new Blob([p.data], { type: p.type });
        const url = URL.createObjectURL(blob);
        const key = p.name.replace(/\.\w+$/, "");
        imagesMap[key] = url;
    });

    return { list, imagesMap };
}
