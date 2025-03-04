import { EventListener } from '@telegram-apps/bridge';
/**
 * True if the component is currently mounted.
 */
export declare const isMounted: import('@telegram-apps/signals').Signal<boolean>;
/**
 * @returns True if the settings button is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Hides the Settings Button.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const hide: () => void;
/**
 * True if the component is currently visible.
 */
export declare const isVisible: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const mount: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * Add a new Settings Button click listener.
 * @param fn - event listener.
 * @returns A function to remove bound listener.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const onClick: import('../../toolkit/withIsSupported.js').WithIsSupported<(fn: EventListener<'settings_button_pressed'>) => VoidFunction>;
/**
 * Removes the Settings Button click listener.
 * @param fn - an event listener.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const offClick: import('../../toolkit/withIsSupported.js').WithIsSupported<(fn: EventListener<'settings_button_pressed'>) => void>;
/**
 * Shows the Settings Button.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const show: () => void;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 *
 * Note that this function does not remove listeners, added via the `onClick` function.
 * @see onClick
 */
export declare function unmount(): void;
