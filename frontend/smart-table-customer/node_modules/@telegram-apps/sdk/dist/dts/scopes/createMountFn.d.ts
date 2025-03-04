import { Signal } from '@telegram-apps/signals';
import { AsyncOptions, CancelablePromise } from '@telegram-apps/bridge';
/**
 * Creates a mount function for a component.
 * @param mount - function mounting the component.
 * @param onMounted - callback which will be called with the mount result.
 * @param mountPromise - signal containing mount promise.
 * @param isMounted - signal containing mount state.
 * @param mountError - signal containing mount error.
 * @param isSupported - signal containing the component support flag.
 */
export declare function createMountFn<T = void>(mount: (options: AsyncOptions) => (T | CancelablePromise<T>), onMounted: (result: T) => void, { isMounting, isMounted, mountError, isSupported, }: {
    isMounted: Signal<boolean>;
    isMounting: Signal<boolean>;
    mountError: Signal<Error | undefined>;
    isSupported?: () => boolean;
}): (options?: AsyncOptions) => CancelablePromise<void>;
