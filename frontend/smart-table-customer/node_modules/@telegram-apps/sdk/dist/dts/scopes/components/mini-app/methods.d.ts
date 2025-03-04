import { BottomBarColor, BackgroundColor } from '@telegram-apps/bridge';
import { isRGB } from '@telegram-apps/transformers';
import { Computed } from '@telegram-apps/signals';
import { GetCssVarNameFn, HeaderColor } from './types.js';
/**
 * True if the Mini App component is supported.
 */
export declare const isSupported: Computed<boolean>;
/**
 * Creates CSS variables connected with the mini app.
 *
 * Default variables:
 * - `--tg-bg-color`
 * - `--tg-header-color`
 * - `--tg-bottom-bar-color`
 *
 * Variables are being automatically updated if theme parameters were changed.
 *
 * @param getCSSVarName - function, returning complete CSS variable name for the specified
 * mini app key.
 * @returns Function to stop updating variables.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_SUPPORTED
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const bindCssVars: (getCSSVarName?: GetCssVarNameFn) => VoidFunction;
/**
 * Closes the Mini App.
 * @param returnBack - Should the client return to the previous activity.
 */
export declare function close(returnBack?: boolean): void;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 *
 * Internally, the function mounts the Theme Params component to work with correctly extracted
 * theme palette values.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const mount: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
/**
 * Informs the Telegram app that the Mini App is ready to be displayed.
 *
 * It is recommended to call this method as early as possible, as soon as all essential
 * interface elements loaded.
 *
 * Once this method is called, the loading placeholder is hidden and the Mini App shown.
 *
 * If the method is not called, the placeholder will be hidden only when the page was fully loaded.
 */
export declare function ready(): void;
/**
 * Updates the background color.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const setBackgroundColor: import('../../toolkit/withIsSupported.js').WithIsSupported<(color: BackgroundColor) => void>;
/**
 * Updates the bottom bar background color.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const setBottomBarColor: import('../../toolkit/withIsSupported.js').WithIsSupported<(color: BottomBarColor) => void>;
/**
 * Updates the header color.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const setHeaderColor: import('../../toolkit/withSupports.js').WithSupports<import('../../toolkit/withIsSupported.js').WithIsSupported<(color: HeaderColor) => void>, {
    color: ["web_app_set_header_color", "color", typeof isRGB];
}>;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 */
export declare function unmount(): void;
