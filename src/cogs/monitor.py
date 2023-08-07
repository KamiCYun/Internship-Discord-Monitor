import discord
from discord import app_commands
from discord.ext import commands, tasks
from discord.ext.commands import Context

from helpers import checks

class Monitor(commands.Cog, name="Monitor"):
    def __init__(self, bot):
        self.bot = bot
        self.jobs = set()
        self.channels = set()

    @commands.hybrid_command(
        name="subscribe", description="Starts sending monitor updates to the channel the command is sent in."
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def subscribe(self, context: Context) -> None:
        """
        Starts sending monitor updates to the channel the command is sent in.

        :param context: The hybrid command context.
        """

        embed = discord.Embed(
                description=f"Channel subscribed to monitor!",
                color=0xD75BF4,
        )
        if context.author.guild_permissions.administrator:
            if context.channel.id not in self.channels:
                self.channels.add(context.channel.id)
            else:
                embed.description = "Channel already subscribed"
        else:
            embed.description = "This command requires administrator"

        await context.send(embed=embed)

    @commands.hybrid_command(
        name="unsubscribe", description="Stops sending monitor updates to the channel the command is sent in."
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def unsubscribe(self, context: Context) -> None:
        """
        Stops sending monitor updates to the channel the command is sent in.

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
                description=f"Channel unsubscribed to monitor!",
                color=0xD75BF4,
        )
        if context.author.guild_permissions.administrator:
            if context.channel.id in self.channels:
                self.channels.remove(context.channel.id)
            else:
                embed.description = "Channel not subscribed"
        else:
            embed.description = "This command requires administrator"

        await context.send(embed=embed)

    @commands.hybrid_command(
        name="start", description="Starts monitoring the github repository for new internships"
    )
    @checks.is_owner()
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def start(self, context: Context) -> None:
        """
        Starts monitoring the github repository for new internships

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            description=f"Monitor Started!",
            color=0xD75BF4,
        )
        try:
            self.monitor_internships.start()
            await context.send(embed=embed)
        except:
            embed.description = "Monitor is already activated"
            await context.send(embed=embed)
    
    @commands.hybrid_command(
        name="stop", description="Stops monitoring the github repository for new internships"
    )
    @checks.is_owner()
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def stop(self, context: Context) -> None:
        """
        Stops monitoring the github repository for new internships

        :param context: The hybrid command context.
        """
        embed = discord.Embed(
            description=f"Monitor Stopped!",
            color=0xD75BF4,
        )
        try:
            await self.monitor_internships.cancel()
            context.send(embed=embed)
        except:
            embed.description = "Monitor is not activated"
            await context.send(embed=embed)

    @tasks.loop(seconds=1)
    async def monitor_internships(self) -> None:
        print(self.channels)


async def setup(bot):
    await bot.add_cog(Monitor(bot))