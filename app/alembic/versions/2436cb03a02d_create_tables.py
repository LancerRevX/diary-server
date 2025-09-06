"""create tables

Revision ID: 2436cb03a02d
Revises:
Create Date: 2025-09-06 15:06:29.667340

"""

from ast import IfExp
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = "2436cb03a02d"
down_revision: Union[str, Sequence[str], None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    op.create_table(
        "users",
        sa.Column("id", sa.BigInteger, primary_key=True),
        if_not_exists=True,
    )
    op.create_table(
        "records",
        sa.Column("id", sa.BigInteger, primary_key=True),
        sa.Column("content", sa.Text, nullable=False),
        if_not_exists=True
    )


def downgrade() -> None:
    op.drop_table("records", if_exists=True)
    op.drop_table("users", if_exists=True)
