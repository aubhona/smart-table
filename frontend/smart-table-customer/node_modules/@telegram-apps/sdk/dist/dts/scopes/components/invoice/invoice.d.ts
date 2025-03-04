import { InvoiceStatus, ExecuteWithPostEvent } from '@telegram-apps/bridge';
/**
 * True if the invoice is currently opened.
 */
export declare const isOpened: import('@telegram-apps/signals').Signal<boolean>;
/**
 * @returns True if the Invoice is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Opens an invoice using its slug.
 * Example of the value: `jd231xxSd1`
 * @param slug - invoice slug.
 * @param options - additional options.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_INVALID_HOSTNAME
 * @throws {TypedError} ERR_INVALID_SLUG
 */
export declare function _open(slug: string, options?: ExecuteWithPostEvent): Promise<InvoiceStatus>;
/**
 * Opens an invoice using its url.
 *
 * The function expects to pass a link in a full format, with the hostname "t.me".
 * Examples:
 * - `https://t.me/$jd231xxSd1`
 * - `https://t.me/invoice/jd231xxSd1`
 * @param url - invoice URL.
 * @param type - value type.
 * @param options - additional options.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_INVALID_HOSTNAME
 * @throws {TypedError} ERR_INVALID_SLUG
 */
export declare function _open(url: string, type: 'url', options?: ExecuteWithPostEvent): Promise<InvoiceStatus>;
/**
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const open: import('../../toolkit/withIsSupported.js').WithIsSupported<typeof _open>;
