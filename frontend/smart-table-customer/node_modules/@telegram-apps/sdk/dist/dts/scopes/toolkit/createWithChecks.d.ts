import { MethodName } from '@telegram-apps/bridge';
import { WithIsSupported } from './withIsSupported.js';
export declare function createWithChecks(isSupportedOrMethod: MethodName | (() => boolean), isMounted: () => boolean): <Fn extends (...args: any[]) => any>(fn: Fn) => WithIsSupported<Fn>;
