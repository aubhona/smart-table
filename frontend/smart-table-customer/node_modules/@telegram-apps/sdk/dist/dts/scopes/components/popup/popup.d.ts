import { OpenOptions } from './types.js';
/**
 * @returns True if the Popup is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * True if a popup is currently opened.
 */
export declare const isOpened: import('@telegram-apps/signals').Signal<boolean>;
/**
 * A method that shows a native popup described by the `params` argument.
 * The promise will be resolved when the popup is closed. Resolved value will have
 * an identifier of pressed button.
 *
 * If a user clicked outside the popup or clicked the top right popup close button, null will be
 * returned.
 *
 * @param options - popup parameters.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_POPUP_INVALID_PARAMS: Invalid title length.
 * @throws {TypedError} ERR_POPUP_INVALID_PARAMS: Invalid message length.
 * @throws {TypedError} ERR_POPUP_INVALID_PARAMS: Invalid buttons length.
 * @throws {TypedError} ERR_POPUP_INVALID_PARAMS: Invalid button id length.
 * @throws {TypedError} ERR_POPUP_INVALID_PARAMS: Invalid text length.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const open: import('../../toolkit/withIsSupported.js').WithIsSupported<(options: OpenOptions) => Promise<string | null>>;
