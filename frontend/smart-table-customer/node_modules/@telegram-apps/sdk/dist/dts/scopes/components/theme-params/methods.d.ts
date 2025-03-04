import { GetCssVarNameFn } from './types.js';
/**
 * Creates CSS variables connected with the current theme parameters.
 *
 * By default, created CSS variables names are following the pattern "--tg-theme-{name}", where
 * {name} is a theme parameters key name converted from camel case to kebab case.
 *
 * Default variables:
 * - `--tg-theme-bg-color`
 * - `--tg-theme-secondary-text-color`
 *
 * Variables are being automatically updated if theme parameters were changed.
 *
 * @param getCSSVarName - function, returning complete CSS variable name for the specified
 * theme parameters key.
 * @returns Function to stop updating variables.
 * @throws TypedError ERR_ALREADY_CALLED
 */
export declare const bindCssVars: (getCSSVarName?: GetCssVarNameFn) => VoidFunction;
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
