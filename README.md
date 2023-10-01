<img src="https://github.com/gabrieljackson/mattermost-plugin-wrangler/blob/master/assets/profile.png?raw=true" width="75" height="75" alt="wrangler">

# Mattermost Wrangler Plugin

![CI](https://github.com/gabrieljackson/mattermost-plugin-wrangler/actions/workflows/ci.yaml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/gabrieljackson/mattermost-plugin-wrangler)](https://goreportcard.com/report/github.com/gabrieljackson/mattermost-plugin-wrangler)

A [Mattermost](https://mattermost.com) plugin for managing messages.

## About

Sometimes Mattermost messages just don't end up in the right place. Wrangler can help with that by providing the following:

  1. Move a single message or thread to a new channel.
  2. Copy a single message or thread to a new channel.
  3. Attach non-threaded messages to a thread.

These functions are designed to quickly bring messages to a place they likely have more relevance in. Example uses include moving a question to a channel where users have direct expertise or to attach a single message to a thread that it obviously was related to.

## Install

1. Go the releases page and download the latest release.
2. On your Mattermost, go to System Console -> Plugin Management and upload it.
3. Configure plugin settings as desired.
4. In order for the plugin to properly recreate messages, ensure the following system console settings are set to true:
    1. [Enable integrations to override usernames](https://docs.mattermost.com/administration/config-settings.html#enable-integrations-to-override-usernames)
    2. [Enable integrations to override profile picture icons](https://docs.mattermost.com/administration/config-settings.html#enable-integrations-to-override-profile-picture-icons)
5. Start using the plugin!

## Questions?

Start by reviewing the [FAQ](#faq) if you have any questions. Feel free to open a GitHub issue if you need additional assistance.

## Commands

Type `/wrangler` for a list of all Wrangler commands.

```
Wrangler Plugin - Slash Command Help

/wrangler move thread [MESSAGE_ID] [CHANNEL_ID]
  Move a given message, along with the thread it belongs to, to a given channel
    - This can be on any channel in any team that you have joined
    - Use the '/wrangler list' commands to get message and channel IDs
    Flags:
      --show-root-message-in-summary   Show the root message in the post-move summary (default true)
      --silent                         Silence all Wrangler summary messages and user DMs when moving the thread


/wrangler copy thread [MESSAGE_ID] [CHANNEL_ID]
  Copy a given message, along with the thread it belongs to, to a given channel
    - This can be on any channel in any team that you have joined
    - Obtain the message ID by running '/wrangler list messages' or via the 'Permalink' message dropdown option (it's the last part of the URL)
    - Obtain the channel ID by running '/wrangler list channels' or via the channel 'View Info' option

/wrangler attach message [MESSAGE_ID_TO_ATTACH] [ROOT_MESSAGE_ID]
  Attach a given message to a thread in the same channel
    - Obtain the message IDs by running '/wrangler list messages' or via the 'Permalink' message dropdown option (it's the last part of the URL)

/wrangler list channels [flags]
  List the IDs of all channels you have joined
    Flags:
      --channel-filter string   A filter value that channel names must contain to be shown on the list
      --team-filter string      A filter value that team names must contain to be shown on the list

/wrangler list messages [flags]
  List the IDs of recent messages in this channel
    Flags:
      --count int         Number of messages to return. Must be between 1 and 100 (default 20)
      --trim-length int   The max character count of messages listed before they are trimmed. Must be between 10 and 500 (default 50)

/wrangler info
  Shows plugin information
```

#### /wrangler move thread

A powerful command that can "move" a message along with its parent thread to a new channel.

Note that the command works by creating new messages in the target channel, but preserves most of the original message metadata. Ordering is kept intact, but the messages contain new timestamps so that channel message history is not altered.

##### Example

A thread that was started in `channel1` is moved to `channel2`.

![channel1](https://user-images.githubusercontent.com/3694686/73672948-d1066380-467b-11ea-9772-097f9fdfcdf0.png)

The thread after being "moved" to `channel2`.

![channel2](https://user-images.githubusercontent.com/3694686/73672959-d499ea80-467b-11ea-97dc-4a2e33c8829e.png)

#### /wrangler copy thread

Similar to the move command, this will duplicate a message or thread and put the copy in another new channel.

#### /wrangler attach message

Attaches a message that is not currently in a thread to an existing message or thread in the same channel.

This is useful for bringing normal messages about a topic into threads that they relate to.

#### /wrangler list channels

Lists channel IDs that you belong to across all teams.

#### /wrangler list channels

Lists recent message IDs from the current channel.

#### /wrangler info

Shows version and commit information for the currently-running plugin build.

## Configuration Options

The following plugin configuration is available:

 - Permitted Wrangler Users: Choose who is allowed to use the Wrangler plugin.
 - Allowed Email Domain: (Optional) When set, users must have an email ending in this domain to use Wrangler. Multiple domains can be specified by separating them with commas.
   - Example: `domain1.com,domain2.net,domain3.org`
 - Enable Wrangler webapp functionality: Enable the work-in-progress Wrangler webapp functionality.
 - Enable Wrangler Command AutoComplete: Control whether command autocomplete is enabled or not. If enabled and Allowed Email Domain is set, then some users will be able to see the Wrangler commands, but will be unable to run them.
 - Max Thread Count Move Size: an optional setting to limit the size of threads that can be moved
 - Enable Moving Threads To Different Teams: Control whether Wrangler is permitted to move message threads from one team to another or not.
 - Enable Moving Threads From Private Channels: Control whether Wrangler is permitted to move message threads from private channels or not.
 - Enable Moving Threads From Direct Message Channels: Control whether Wrangler is permitted to move message threads from direct message channels or not.
 - Enable Moving Threads From Group Message Channels: Control whether Wrangler is permitted to move message threads from group message channels or not.
 - Message customization: Various customization options are available to tailor the direct messages that are sent from Wrangler.

## FAQ

Q: I would very much like some UI please.

A: That isn't a question, but I hear you. The Wrangler UI can be enabled in plugin settings in the system console.

---

Q: Why don't I see any command autocomplete options when I type `/wrangler`?

A: Command autocomplete can also be enabled in plugin settings in the system console. Keep in mind that all users can see autocomplete prompts even if they don't have permissions to use Wrangler.

---

Q: It would be awesome if Wrangler could do this other thing! Is that coming any time soon?

A: Please open a GitHub issue and I will see if we can implement it. Keep in mind that security and stability is always prioritized when adding features that may have edge cases where message management could have undesired outcomes.

---

Q: Why do I see `(message deleted)` when I use Wrangler?

A: Wrangler often performs message management tasks by recreating messages and deleting the originals. This results in a fairly seamless experience, but does have some side effects that can't be resolved without some core changes to Mattermost itself. The `(message deleted)` system posts are the most common example of this. Ideally, future changes to Mattermost could allow for direct message movement which would provide a more seamless and efficient user experience. For now, this is the best we can do.

---

Q: When I run `/wranger attach message` it seems like the attached message is out of order?

A: When attaching a message, it's necessary to create a new post in the thread which triggers the default behavior of Mattermost to show the message at the bottom of the channel. The message has been attached to the thread with the correct timestamp of when it was originally posted though, so simply reloading the channel will resolve the out-of-order behavior you are initially experiencing. This is also something I would like to improve in the future if possible. (Note that this behavior was changed for Wrangler after v0.3.0 was cut)
