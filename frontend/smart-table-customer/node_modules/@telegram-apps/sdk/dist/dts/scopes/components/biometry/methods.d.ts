import { CancelablePromise, BiometryTokenUpdateStatus, BiometryAuthRequestStatus } from '@telegram-apps/bridge';
import { AuthenticateOptions, RequestAccessOptions, UpdateTokenOptions } from './types.js';
/**
 * @returns True if the biometry manager is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Attempts to authenticate a user using biometrics and fetch a previously stored
 * secure token.
 * @param options - method options.
 * @since 7.2
 * @returns Token from the local secure storage saved previously or undefined.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_AVAILABLE
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const authenticate: (options?: AuthenticateOptions) => CancelablePromise<{
    /**
     * Authentication status.
     */
    status: BiometryAuthRequestStatus;
    /**
     * Token from the local secure storage saved previously.
     */
    token?: string;
}>;
/**
 * Opens the biometric access settings for bots. Useful when you need to request biometrics
 * access to users who haven't granted it yet.
 *
 * _Note that this method can be called only in response to user interaction with the Mini App
 * interface (e.g. a click inside the Mini App or on the main button)_.
 * @since 7.2
 */
export declare const openSettings: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * Requests permission to use biometrics.
 * @since 7.2
 * @returns Promise with true, if access was granted.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_AVAILABLE
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const requestAccess: (options?: RequestAccessOptions) => CancelablePromise<boolean>;
/**
 * Mounts the component.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const mount: (options?: import('@telegram-apps/bridge').AsyncOptions | undefined) => CancelablePromise<void>;
/**
 * Unmounts the component.
 */
export declare function unmount(): void;
/**
 * Updates the biometric token in a secure storage on the device.
 * @since 7.2
 * @returns Promise with `true`, if token was updated.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const updateToken: (options?: UpdateTokenOptions) => CancelablePromise<BiometryTokenUpdateStatus>;
