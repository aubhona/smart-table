import { ExecuteWithOptions, CancelablePromise } from '@telegram-apps/bridge';
import { State } from './types.js';
/**
 * Requests biometry information.
 * @param options - additional execution options.
 */
export declare const requestBiometry: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => CancelablePromise<State>>;
