import { GetCSSVarNameFn } from './static.js';
/**
 * Creates CSS variables connected with the current viewport.
 *
 * By default, created CSS variables names are following the pattern "--tg-theme-{name}", where
 * {name} is a theme parameters key name converted from camel case to kebab case.
 *
 * Default variables:
 * - `--tg-viewport-height`
 * - `--tg-viewport-width`
 * - `--tg-viewport-stable-height`
 *
 * Variables are being automatically updated if viewport was changed.
 *
 * @param getCSSVarName - function, returning complete CSS variable name for the specified
 * viewport property.
 * @returns Function to stop updating variables.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_MOUNTED
 */
export declare const bindCssVars: (getCSSVarName?: GetCSSVarNameFn) => VoidFunction;
/**
 * A method that expands the Mini App to the maximum available height. To find out if the Mini
 * App is expanded to the maximum height, refer to the value of the `isExpanded`.
 * @see isExpanded
 */
export declare function expand(): void;
/**
 * Mounts the component.
 *
 * This function restores the component state and is automatically saving it in the local storage
 * if it changed.
 */
export declare const mount: (options?: import('@telegram-apps/bridge').AsyncOptions | undefined) => import('@telegram-apps/bridge').CancelablePromise<void>;
/**
 * Unmounts the component, removing the listener, saving the component state in the local storage.
 */
export declare function unmount(): void;
