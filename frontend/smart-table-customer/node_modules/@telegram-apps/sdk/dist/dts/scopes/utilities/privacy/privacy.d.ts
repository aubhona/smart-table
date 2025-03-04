import { CancelablePromise, PhoneRequestedStatus, WriteAccessRequestedStatus, ExecuteWithOptions } from '@telegram-apps/bridge';
/**
 * Requested contact information.
 */
export interface RequestedContact {
    contact: {
        userId: number;
        phoneNumber: string;
        firstName: string;
        lastName?: string;
    };
    authDate: Date;
    hash: string;
}
/**
 * True if phone access is currently being requested.
 */
export declare const isRequestingPhoneAccess: import('@telegram-apps/signals').Signal<boolean>;
/**
 * True if write access is currently being requested.
 */
export declare const isRequestingWriteAccess: import('@telegram-apps/signals').Signal<boolean>;
/**
 * Requests current user contact information. In contrary to requestPhoneAccess, this method
 * returns promise with contact information that rejects in case, user denied access, or request
 * failed.
 * @param options - additional options.
 * @throws {TypedError} ERR_ACCESS_DENIED
 * @throws {TypedError} ERR_CUSTOM_METHOD_ERR_RESPONSE
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const requestContact: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => CancelablePromise<RequestedContact>>;
/**
 * Requests current user phone access. Method returns promise, which resolves
 * status of the request. In case, user accepted the request, Mini App bot will receive
 * the according notification.
 *
 * To obtain the retrieved information instead, utilize the `requestContact` method.
 * @param options - additional options.
 * @see requestContact
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const requestPhoneAccess: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => Promise<PhoneRequestedStatus>>;
/**
 * Requests write message access to the current user.
 * @param options - additional options.
 * @throws {TypedError} ERR_ALREADY_CALLED
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const requestWriteAccess: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => Promise<WriteAccessRequestedStatus>>;
