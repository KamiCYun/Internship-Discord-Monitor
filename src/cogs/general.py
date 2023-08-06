import platform

import discord
from discord import app_commands
from discord.ext import commands
from discord.ext.commands import Context


class General(commands.Cog, name="general"):
    def __init__(self, bot):
        self.bot = bot

    @commands.hybrid_command(
        name="help", description="List all commands the bot has loaded."
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def help(self, context: Context) -> None:
        prefix = self.bot.config["prefix"]
        embed = discord.Embed(
            title="Help", description="List of available commands:", color=0x9C84EF
        )
        for i in self.bot.cogs:
            cog = self.bot.get_cog(i.lower())
            commands = cog.get_commands()
            data = []
            for command in commands:
                description = command.description.partition("\n")[0]
                data.append(f"{prefix}{command.name} - {description}")
            help_text = "\n".join(data)
            embed.add_field(
                name=i.capitalize(), value=f"```{help_text}```", inline=False
            )
        await context.send(embed=embed)

    @commands.hybrid_command(
        name="botinfo",
        description="Get some useful (or not) information about the bot.",
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def botinfo(self, context: Context) -> None:
        """
        Get some useful (or not) information about the bot.

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            description="Sends real-time updates from the [Pitt CSC Internship Repo](https://github.com/SimplifyJobs/Summer2024-Internships/tree/dev)",
            color=0x9C84EF,
        )
        embed.set_author(name="Bot Information")
        embed.add_field(name="Developer:", value="kamicy", inline=True)
        embed.add_field(
            name="Python Version:", value=f"{platform.python_version()}", inline=True
        )
        embed.add_field(
            name="Prefix:",
            value=f"/ (Slash Commands) or {self.bot.config['prefix']} for normal commands",
            inline=False,
        )
        embed.set_footer(text=f"Requested by {context.author}")
        await context.send(embed=embed)

    @commands.hybrid_command(
        name="ping",
        description="Check if the bot is alive.",
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def ping(self, context: Context) -> None:
        """
        Check if the bot is alive.

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            title="🏓 Pong!",
            description=f"The bot latency is {round(self.bot.latency * 1000)}ms.",
            color=0x9C84EF,
        )
        await context.send(embed=embed)

    @commands.hybrid_command(
        name="invite",
        description="Get the invite link of the bot to be able to invite it.",
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def invite(self, context: Context) -> None:
        """
        Get the invite link of the bot to be able to invite it.

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            description=f"Invite me by clicking [here](https://discordapp.com/oauth2/authorize?&client_id={self.bot.config['application_id']}&scope=bot+applications.commands&permissions={self.bot.config['permissions']}).",
            color=0xD75BF4,
        )
        try:
            # To know what permissions to give to your bot, please see here: https://discordapi.com/permissions.html and remember to not give Administrator permissions.
            await context.author.send(embed=embed)
            await context.send("Check your DMs for the Link!")
        except discord.Forbidden:
            await context.send(embed=embed)

    @commands.hybrid_command(
        name="github",
        description="Get the link to the github repository of the bot.",
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def server(self, context: Context) -> None:
        """
        Get the link to the github repository of the bot.

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            description=f"Check out and star the github repository for the bot by clicking [here](https://github.com/KamiCYun/Internship-Discord-Monitor).",
            color=0xD75BF4,
        )
        try:
            await context.send(embed=embed)
        except discord.Forbidden:
            await context.send(embed=embed)


async def setup(bot):
    await bot.add_cog(General(bot))