import { CancelablePromise, ExecuteWithOptions, SwitchInlineQueryChatType, ExecuteWithPostEvent } from '@telegram-apps/bridge';
export interface ShareStoryOptions extends ExecuteWithPostEvent {
    /**
     * The caption to be added to the media.
     * 0-200 characters for regular users and 0-2048 characters for premium subscribers.
     * @see https://telegram.org/faq_premium#telegram-premium
     */
    text?: string;
    /**
     * An object that describes a widget link to be included in the story.
     * Note that only premium subscribers can post stories with links.
     * @see https://telegram.org/faq_premium#telegram-premium
     */
    widgetLink?: {
        /**
         * The URL to be included in the story.
         */
        url: string;
        /**
         * The name to be displayed for the widget link, 0-48 characters.
         */
        name?: string;
    };
}
/**
 * Reads a text from the clipboard and returns a string or null. null is returned
 * in cases:
 * - A value in the clipboard is not a text.
 * - Access to the clipboard is not granted.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const readTextFromClipboard: import('../../toolkit/withIsSupported.js').WithIsSupported<(options?: ExecuteWithOptions) => CancelablePromise<string | null>>;
/**
 * A method used to send data to the bot.
 *
 * When this method called, a service message sent to the bot containing the data of the length
 * up to 4096 bytes, and the Mini App closed.
 *
 * See the field `web_app_data` in the class [Message](https://core.telegram.org/bots/api#message).
 *
 * This method is only available for Mini Apps launched via a Keyboard button.
 * @param data - data to send to bot.
 * @throws {TypedError} ERR_DATA_INVALID_SIZE
 */
export declare function sendData(data: string): void;
/**
 * A method that opens the native story editor.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const shareStory: import('../../toolkit/withIsSupported.js').WithIsSupported<(mediaUrl: string, options?: ShareStoryOptions) => void>;
/**
 * Inserts the bot's username and the specified inline query in the current chat's input field.
 * Query may be empty, in which case only the bot's username will be inserted. The client prompts
 * the user to choose a specific chat, then opens that chat and inserts the bot's username and
 * the specified inline query in the input field.
 * @param query - text which should be inserted in the input after the current bot name. Max
 * length is 256 symbols.
 * @param chatTypes - List of chat types which could be chosen to send the message. Could be
 * empty list.
 * @throws {TypedError} ERR_NOT_SUPPORTED
 */
export declare const switchInlineQuery: import('../../toolkit/withIsSupported.js').WithIsSupported<(query: string, chatTypes?: SwitchInlineQueryChatType[]) => void>;
