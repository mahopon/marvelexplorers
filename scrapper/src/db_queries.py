import psycopg
from psycopg import extras
from classes import Character
from . import database as db

def insertCharacters(chars: list[Character]):
    with db.PGPool().connect() as conn:
        with conn.cursor() as cur:
            extras.execute_values(
                cur,
                "INSERT INTO Characters VALUES %s",
                [repr(char) for char in chars]
            )