-- User table---
CREATE TABLE IF NOT EXISTS  "Users" (
    "ID" INTEGER PRIMARY KEY AUTOINCREMENT,
    "FirstName" TEXT NOT NULL,
    "LastName" TEXT NOT NULL,
    "CreatedDate" DATETIME NOT NULL,
    "IsActive" BOOLEAN
);
--------------------------------
-- Message table----------------

CREATE TABLE IF NOT EXISTS  "Messages"(
    "ID"              INTEGER PRIMARY KEY AUTOINCREMENT,
	"Subject"         TEXT NOT NULL,
	"MessageBody"     TEXT NOT NULL,
	"SenderID"       INTEGER NOT NULL,
	"ParentMessageID" INTEGER,
	"CreatedDate"     DATETIME NOT NULL,
	"ExpiryDate"      DATETIME,
    FOREIGN KEY("SenderID") REFERENCES "Users(ID)"
);
----------------------------------

--- Group table -----------------

CREATE TABLE IF NOT EXISTS  "Groups"(
   	"ID"          INTEGER PRIMARY KEY AUTOINCREMENT,
	"Name"        TEXT NOT NULL,
	"CreatedDate" DATETIME NOT NULL,
	"IsActive"    BOOLEAN
);

---------------------------------
----User Groups-----------------

CREATE TABLE IF NOT EXISTS  "UserGroup"(
	"ID"          INTEGER PRIMARY KEY AUTOINCREMENT,
	"UserID"      INTEGER NOT NULL,
	"GroupID"     INTEGER NOT NULL,
    "AdminUserID" INTEGER NOT NULL,  
	"CreatedDate" DATETIME NOT NULL,
	"IsActive"    BOOLEAN,
    FOREIGN KEY("UserID") REFERENCES "Users(ID)",
    FOREIGN KEY("GroupID") REFERENCES "Groups(ID)",
    FOREIGN KEY("AdminUserID") REFERENCES "Users(ID)"
);
-----------------------------------
----Message REcipient-------------

CREATE TABLE IF NOT EXISTS   "MessageRecipient"(
	"ID"               INTEGER PRIMARY KEY AUTOINCREMENT,
	"RecipientID"      INTEGER NOT NULL,
	"RecipientGroupID" INTEGER NOT NULL,
	"MessageID"        INTEGER NOT NULL,
	"IsRead"           BOOLEAN,
    FOREIGN KEY("RecipientID") REFERENCES "Users(ID)",
    FOREIGN KEY("RecipientGroupID") REFERENCES "Groups(ID)",
    FOREIGN KEY("MessageID") REFERENCES "Messages(ID)"
);
----------------------------------