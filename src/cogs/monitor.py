import discord
from discord import app_commands
from discord.ext import commands, tasks
from discord.ext.commands import Context
import re
import requests
from hashlib import sha256
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
        headers = {
            'authority': 'raw.githubusercontent.com',
            'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7',
            'accept-language': 'en-US,en;q=0.9',
            'cache-control': 'max-age=0',
            'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36',
        }

        response = requests.get('https://raw.githubusercontent.com/KamiCYun/Internship-Discord-Monitor/main/test.md', headers=headers)
        jobs = re.findall("\|\s\*{2}.+\*{2}\s\|\s.+\s\|\n", response.text)

        # make this better once u get some sleep :)
        for job in jobs:
            fields = job.split(" | ")
            hash_str = fields[0] + fields[1] + fields[2] + fields[4]
            job_hash = sha256(hash_str.encode('utf-8')).hexdigest()
            if job_hash not in self.jobs:
                title = re.findall("\*\*\[(.+)\].*\*\*", fields[0])
                if len(title) == 0:
                    title = re.findall("\*\*(.+)\*\*", fields[0])
                title = title[0]

                position = fields[1]
                location = fields[2]
                url = re.findall("ref=\".*\">", fields[3])
                if len(url) == 0:
                    url = "Closed"
                else:
                    url = url[0]
                date = fields[4].replace(" |", "")

                embed = discord.Embed(
                    title=f"New Job | {title}",
                    description=position,
                    color=0xD75BF4,
                )

                embed.add_field(
                    name="Location", value=location, inline=False
                )
                embed.add_field(
                    name="Application", value=url, inline=False
                )
                embed.add_field(
                    name="Date Posted", value=date, inline=False
                )

                for channel in self.channels:
                    c = self.bot.get_channel(channel)
                    await c.send(embed=embed)


async def setup(bot):
    await bot.add_cog(Monitor(bot))