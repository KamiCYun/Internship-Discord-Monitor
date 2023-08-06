import discord
from discord import app_commands
from discord.ext import commands, tasks
from discord.ext.commands import Context


class Monitor(commands.Cog, name="Monitor"):
    def __init__(self, bot):
        self.test.start()
        print('hi')
        self.bot = bot
        self.jobs = set()
        self.channels = set()

    @commands.hybrid_command(
        name="add", description="Starts sending monitor updates to the channel the command is sent in."
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def help(self, context: Context) -> None:
        """
        Starts sending monitor updates to the channel the command is sent in.

        :param context: The hybrid command context.
        """
        self.channels.add(context)

    @commands.hybrid_command(
        name="remove", description="Stops sending monitor updates to the channel the command is sent in."
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def help(self, context: Context) -> None:
        """
        Stops sending monitor updates to the channel the command is sent in.

        :param context: The hybrid command context.
        """
        self.channels.add(context)

    @commands.hybrid_command(
        name="start", description="Starts monitoring the github repository for new internships"
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def help(self, context: Context) -> None:
        """
        Starts monitoring the github repository for new internships

        :param context: The hybrid command context.
        """
        self.channels.add(context)
    
    @commands.hybrid_command(
        name="stop", description="Stops monitoring the github repository for new internships"
    )
    @app_commands.guilds(discord.Object(id=1106017357643644948)) 
    async def help(self, context: Context) -> None:
        """
        Stops monitoring the github repository for new internships

        :param context: The hybrid command context.
        """
        self.channels.add(context)

    @tasks.loop(seconds=1)
    async def test(self) -> None:
        print("hihi")


async def setup(bot):
    await bot.add_cog(Monitor(bot))