import { CancelablePromise, ExecuteWithOptions } from '@telegram-apps/bridge';
interface OpenSharedOptions extends ExecuteWithOptions {
    /**
     * Title to be displayed in the scanner.
     */
    text?: string;
}
/**
 * @returns True if the QR scanner is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Closes the scanner.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const close: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * True if the scanner is currently opened.
 */
export declare const isOpened: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Opens the scanner and returns a promise which will be resolved with the QR content if the
 * `capture` function returned true.
 *
 * Promise may also be resolved to null if the scanner was closed.
 * @param options - method options.
 * @returns A promise with QR content presented as string or undefined if the scanner was closed.
 * @throws {TypedError} ERR_ALREADY_CALLED
 */
declare function _open(options?: OpenSharedOptions & {
    /**
     * Function, which should return true if a scanned QR should be captured.
     * @param qr - scanned QR content.
     */
    capture?: (qr: string) => boolean;
}): CancelablePromise<string | undefined>;
/**
 * Opens the scanner and calls the `onCaptured` function each time, a QR was scanned.
 *
 * The method does not return anything and expects the scanner to be closed externally by a user
 * or via the `close` method.
 * @param options - method options.
 * @throws {TypedError} ERR_ALREADY_CALLED
 */
declare function _open(options: OpenSharedOptions & {
    /**
     * Function which will be called if some QR code was scanned.
     * @param qr - scanned QR content.
     */
    onCaptured: (qr: string) => void;
}): CancelablePromise<void>;
/**
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const open: import('../../toolkit/withIsSupported.js').WithIsSupported<typeof _open>;
export {};
