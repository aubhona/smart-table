import { EventListener } from '@telegram-apps/bridge';
import { State } from './types.js';
/**
 * @returns True if the Secondary Button is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const mount: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * Adds a new main button click listener.
 * @param fn - event listener.
 * @returns A function to remove bound listener.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const onClick: import('../../toolkit/withIsSupported.js').WithIsSupported<(fn: EventListener<'secondary_button_pressed'>) => VoidFunction>;
/**
 * Removes the main button click listener.
 * @param fn - an event listener.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const offClick: import('../../toolkit/withIsSupported.js').WithIsSupported<(fn: EventListener<'secondary_button_pressed'>) => void>;
/**
 * Updates the main button state.
 * @param updates - state changes to perform.
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const setParams: (updates: Partial<State>) => void;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 *
 * Note that this function does not remove listeners, added via the `onClick` function.
 * @see onClick
 */
export declare function unmount(): void;
