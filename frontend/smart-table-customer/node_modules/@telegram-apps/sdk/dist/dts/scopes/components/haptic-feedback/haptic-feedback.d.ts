import { ImpactHapticFeedbackStyle, NotificationHapticFeedbackType } from '@telegram-apps/bridge';
/**
 * @returns True if the Haptic Feedback is supported.
 */
export declare const isSupported: import('@telegram-apps/signals').Computed<boolean>;
/**
 * A method tells that an impact occurred. The Telegram app may play the appropriate haptics based
 * on style value passed.
 * @param style - impact style.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const impactOccurred: import('../../toolkit/withIsSupported.js').WithIsSupported<(style: ImpactHapticFeedbackStyle) => void>;
/**
 * A method tells that a task or action has succeeded, failed, or produced a warning. The Telegram
 * app may play the appropriate haptics based on type value passed.
 * @param type - notification type.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const notificationOccurred: import('../../toolkit/withIsSupported.js').WithIsSupported<(type: NotificationHapticFeedbackType) => void>;
/**
 * A method tells that the user has changed a selection. The Telegram app may play the
 * appropriate haptics.
 *
 * Do not use this feedback when the user makes or confirms a selection; use it only when the
 * selection changes.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const selectionChanged: import('../../toolkit/withIsSupported.js').WithIsSupported<() => void>;
