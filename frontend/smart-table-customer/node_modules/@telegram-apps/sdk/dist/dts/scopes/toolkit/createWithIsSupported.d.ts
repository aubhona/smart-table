import { MethodName } from '@telegram-apps/bridge';
import { WithIsSupported } from './withIsSupported.js';
export declare function createWithIsSupported(isSupportedOrMethod: MethodName | (() => boolean)): <Fn extends (...args: any[]) => any>(fn: Fn) => WithIsSupported<Fn>;
