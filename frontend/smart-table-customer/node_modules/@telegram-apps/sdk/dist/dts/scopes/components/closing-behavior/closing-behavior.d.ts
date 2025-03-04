/**
 * True if the component is currently mounted.
 */
export declare const isMounted: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Disables the confirmation dialog when closing the Mini App.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const disableConfirmation: () => void;
/**
 * True if the confirmation dialog should be shown while the user is trying to close the Mini App.
 */
export declare const isConfirmationEnabled: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Enables the confirmation dialog when closing the Mini App.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const enableConfirmation: () => void;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 */
export declare function mount(): void;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 */
export declare function unmount(): void;
