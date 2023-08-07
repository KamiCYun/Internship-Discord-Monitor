import json
import os
from typing import Callable, TypeVar

from discord.ext import commands

from exceptions import *

T = TypeVar("T")


def is_owner() -> Callable[[T], T]:
    """
    This is a custom check to see if the user executing the command is an owner of the bot.
    """

    async def predicate(context: commands.Context) -> bool:
        with open(
            f"{os.path.realpath(os.path.dirname(__file__))}/../../config.json"
        ) as file:
            data = json.load(file)
        if context.author.id not in data["owners"]:
            raise UserNotOwner
        return True

    return commands.check(predicate)