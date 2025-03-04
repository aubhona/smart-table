import { CancelablePromise, ExecuteWithOptions } from '@telegram-apps/bridge';
/**
 * @returns True if the Cloud Storage is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Deletes specified key or keys from the cloud storage.
 * @param keyOrKeys - key or keys to delete.
 * @param options - request execution options.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const deleteItem: import('../../toolkit/withIsSupported.js').WithIsSupported<(keyOrKeys: string | string[], options?: ExecuteWithOptions) => CancelablePromise<void>>;
/**
 * @param keys - keys list.
 * @param options - request execution options.
 * @returns Map, where a key is one of the specified in the `keys` argument, and a value is
 * a corresponding storage value.
 */
declare function _getItem<K extends string>(keys: K[], options?: ExecuteWithOptions): CancelablePromise<Record<K, string>>;
/**
 * @param key - cloud storage key.
 * @param options - request execution options.
 * @return Value of the specified key. If the key was not created previously, the function
 * will return an empty string.
 */
declare function _getItem(key: string, options?: ExecuteWithOptions): CancelablePromise<string>;
/**
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const getItem: import('../../toolkit/withIsSupported.js').WithIsSupported<typeof _getItem>;
/**
 * Returns a list of all keys presented in the cloud storage.
 * @param options - request execution options.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const getKeys: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => CancelablePromise<string[]>>;
/**
 * Saves specified value by key.
 * @param key - storage key.
 * @param value - storage value.
 * @param options - request execution options.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const setItem: import('../../toolkit/withIsSupported.js').WithIsSupported<(key: string, value: string, options?: ExecuteWithOptions) => CancelablePromise<void>>;
export {};
