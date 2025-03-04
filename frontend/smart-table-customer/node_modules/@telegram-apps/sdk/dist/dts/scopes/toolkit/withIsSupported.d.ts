import { MethodName } from '@telegram-apps/bridge';
export type WithIsSupported<F extends (...args: any) => any> = F & {
    /**
     * @returns True, if this function is supported.
     */
    isSupported(): boolean;
};
/**
 * Adds the "isSupported" method to the passed function returning true if the wrapped
 * function is supported.
 *
 * Also, if the "isSupported" method returned false during the function execution,
 * the ERR_NOT_SUPPORTED error will be thrown.
 * @param fn - function to extend.
 * @param isSupported - function to check if the function is supported.
 */
export declare function withIsSupported<F extends (...args: any) => any>(fn: F, isSupported: () => boolean): WithIsSupported<F>;
/**
 * Adds the "isSupported" method to the passed function returning true, if the current Mini Apps
 * version supports passed Mini Apps method.
 *
 * Also, if the "isSupported" method returned false during the function execution,
 * the ERR_NOT_SUPPORTED error will be thrown.
 * @param fn - function to extend.
 * @param method - Mini Apps method.
 */
export declare function withIsSupported<F extends (...args: any) => any>(fn: F, method: MethodName): WithIsSupported<F>;
