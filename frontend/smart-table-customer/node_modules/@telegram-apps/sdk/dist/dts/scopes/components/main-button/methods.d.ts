import { EventListener } from '@telegram-apps/bridge';
import { State } from './types.js';
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 */
export declare function mount(): void;
/**
 * Adds a new main button click listener.
 * @param fn - event listener.
 * @returns A function to remove bound listener.
 */
export declare function onClick(fn: EventListener<'main_button_pressed'>): VoidFunction;
/**
 * Removes the main button click listener.
 * @param fn - an event listener.
 */
export declare function offClick(fn: EventListener<'main_button_pressed'>): void;
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
