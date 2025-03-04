import { ExecuteWithOptions, CancelablePromise } from '@telegram-apps/bridge';
export interface RequestResult {
    height: number;
    isExpanded: boolean;
    isStable: boolean;
    width: number;
}
/**
 * Requests viewport actual information from the Telegram application.
 * @param options - request options.
 */
export declare function request(options?: ExecuteWithOptions): CancelablePromise<RequestResult>;
export type * from './types.js';
