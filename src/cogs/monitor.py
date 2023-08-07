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
        self.first_run = True

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

    @tasks.loop(seconds=10)
    async def monitor_internships(self) -> None:
        print(self.jobs)
        headers = {
            'authority': 'raw.githubusercontent.com',
            'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7',
            'accept-language': 'en-US,en;q=0.9',
            'cache-control': 'max-age=0',
            'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36',
        }

        response = requests.get('https://raw.githubusercontent.com/KamiCYun/Internship-Discord-Monitor/main/src/test.md', headers=headers)
        jobs = re.findall("\|\s?\*{2}.+\*{2}\s?\|\s.+\s\|\n", response.text)

        for job in jobs:
            job_hash = sha256(job.encode('utf-8')).hexdigest()
            
            if job_hash not in self.jobs:
                self.jobs.add(job_hash)
                fields = job.split(" | ")
                title_match = re.search(r'\*{2}(.+?)\*{2}', fields[0])
                title = title_match.group(1) if title_match else "No Title"

                if "[" in title:
                    pattern = r'\[(.*?)\]'
                    matches = re.findall(pattern, title)
                    title = ' '.join(matches)         

                position, location, url = fields[1], fields[2].replace("<br/>",""), fields[3]
                url_match = re.search(r'ref="(.*?)">', url)
                url = f"[Click Me!]({url_match.group(1)})" if url_match else "Closed"
                
                date = fields[4].replace(" |", "")

                if self.first_run == False:
                    embed = discord.Embed(
                        title=f"Internship Update | {title}",
                        color=0xD75BF4,
                    )

                    fields = [
                        ("Position", position, True),
                        ("Location", location, True),
                        ("Date Posted", date, True),
                        ("Application", url, False),
                    ]

                    for name, value, inline in fields:
                        embed.add_field(name=name, value=value, inline=inline)

                    for channel in self.channels:
                        c = self.bot.get_channel(channel)
                        await c.send(embed=embed)
        self.first_run = False

async def setup(bot):
    await bot.add_cog(Monitor(bot))