import { OpenLinkBrowser } from '@telegram-apps/bridge';
export interface OpenLinkOptions {
    /**
     * Attempts to use the instant view mode.
     */
    tryInstantView?: boolean;
    /**
     * A preferred browser to open the link in.
     */
    tryBrowser?: OpenLinkBrowser;
}
/**
 * Opens a link.
 *
 * The Mini App will not be closed.
 *
 * Note that this method can be called only in response to the user
 * interaction with the Mini App interface (e.g. click inside the Mini App or on the main button).
 * @param url - URL to be opened.
 * @param options - additional options.
 */
export declare function openLink(url: string, options?: OpenLinkOptions): void;
/**
 * Opens a Telegram link inside Telegram app. The Mini App will be closed. It expects passing
 * links in full format, with hostname "t.me".
 * @param url - URL to be opened.
 * @throws {TypedError} ERR_INVALID_HOSTNAME
 */
export declare function openTelegramLink(url: string): void;
/**
 * Shares specified URL with the passed to the chats, selected by user. After being called,
 * it closes the mini application.
 *
 * This method uses Telegram's Share Links.
 * @param url - URL to share.
 * @param text - text to append after the URL.
 * @see https://core.telegram.org/api/links#share-links
 * @see https://core.telegram.org/widgets/share#custom-buttons
 */
export declare function shareURL(url: string, text?: string): void;
