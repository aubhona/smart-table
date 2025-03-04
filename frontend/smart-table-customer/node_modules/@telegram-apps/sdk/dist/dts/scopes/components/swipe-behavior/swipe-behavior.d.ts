/**
 * True if the component is currently mounted.
 */
export declare const isMounted: import('@telegram-apps/signals').Signal<boolean>;
/**
 * @returns True if the Swipe Behavior is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Disables vertical swipes.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const disableVertical: () => void;
/**
 * Enables vertical swipes.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const enableVertical: () => void;
/**
 * True if vertical swipes are enabled.
 */
export declare const isVerticalEnabled: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const mount: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 */
export declare function unmount(): void;
