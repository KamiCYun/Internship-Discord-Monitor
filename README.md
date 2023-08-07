# Internship Discord Monitor
A discord monitor that sends new application postings from [Pitt CSC Internships Repo](https://github.com/SimplifyJobs/Summer2024-Internships/tree/dev) in realtime. Discord bot boilerplate code found [here](https://github.com/kkrypt0nn/Python-Discord-Bot-Template/tree/main). Show some love and **star** this repository!

I'm a broke college student so any donations appreciated:
Venmo: @CYun3

## How it works
The ```?subscribe``` and ```?unsubscribe``` commands are used to add/remove a channel to the monitor's output list.

The ```?start``` and ```?stop``` commands are used to start/stop the periodic monitoring of the internship repository.


## Cache bypass method

GitHub employs the Fastly CDN for document delivery, utilizing edge caching to optimize response times and alleviate server load. However, this poses a challenge for real-time GitHub monitoring, as cached content only updates upon expiration, potentially causing delays of hours for repository updates to propagate.

To address this, the solution involves prompting the Fastly CDN to deliver content directly from the origin server. A "HIT" indicates cached content, while a "MISS" denotes content fetched from the origin server. To achieve this, requests must be manipulated to appear to be requesting unique documents, even when referencing the same document. This is effective due to the fact that these "unique" documents do not exist on the edge cache yet.

One effective approach I discovered is to capitalize letters randomly within the URL path. The CDN's case-sensitive URL checks treat the document as distinct due to varying capitalization, consistently triggering "MISS" responses. This ensures timely updates by forcing content to be served from the origin server 100% of the time.

## Disclaimer

Slash commands can take some time to get registered globally, so if you want to test a command you should use
the `@app_commands.guilds()` decorator so that it gets registered instantly. Example:

```py
@commands.hybrid_command(
  name="command",
  description="Command description",
)
@app_commands.guilds(discord.Object(id=GUILD_ID)) # Place your guild ID here
```

## How to download it

* Clone/Download the repository
    * To clone it and get the updates you can definitely use the command
      `git clone`
* Create a discord bot [here](https://discord.com/developers/applications)
* Get your bot token
* Invite your bot on servers using the following invite:
  https://discord.com/oauth2/authorize?&client_id=YOUR_APPLICATION_ID_HERE&scope=bot+applications.commands&permissions=PERMISSIONS (
  Replace `YOUR_APPLICATION_ID_HERE` with the application ID and replace `PERMISSIONS` with the required permissions
  your bot needs that it can be get at the bottom of a this
  page https://discord.com/developers/applications/YOUR_APPLICATION_ID_HERE/bot)

  ## How to set up

To set up the bot I made it as simple as possible. I now created a [config.json](config.json) file where you can put the
needed things to edit.

Here is an explanation of what everything is:

| Variable                  | What it is                                                            |
| ------------------------- | ----------------------------------------------------------------------|
| YOUR_BOT_PREFIX_HERE      | The prefix you want to use for normal commands                        |
| YOUR_BOT_TOKEN_HERE       | The token of your bot                                                 |
| YOUR_BOT_PERMISSIONS_HERE | The permissions integer your bot needs when it gets invited           |
| YOUR_APPLICATION_ID_HERE  | The application ID of your bot                                        |
| OWNERS                    | The user ID of all the bot owners                                     |


## How to start

To start the bot you simply need to launch, either your terminal (Linux, Mac & Windows), or your Command Prompt (
Windows)
.

Before running the bot you will need to install all the requirements with this command:

```
python -m pip install -r requirements.txt
```

After that you can start it with

```
python bot.py
```

> **Note** You may need to replace `python` with `py`, `python3`, `python3.11`, etc. depending on what Python versions you have installed on the machine.
