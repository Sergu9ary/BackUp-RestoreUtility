# My Backup Utility

## Overview

This utility allows you to create **full** and **incremental** backups of a directory and restore them at a later time. The backup system works with two main directories:

- `/work`: This is the directory containing files and subdirectories that the user is actively working on.
- `/backup`: This directory stores backup copies, organized in subdirectories named with the format `YYYY-MM-DD_HH-MM-SS`. These subdirectories contain either full or incremental backups of the `/work` directory.

The utility provides two main commands:
1. **Backup**: Creates a full or incremental backup of the `/work` directory into the `/backup` directory.
2. **Restore**: Restores a backup from the `/backup` directory to the `/work` directory.

### Backup Types:
- **Full Backup**: Creates a complete snapshot of the `/work` directory, copying all files and subdirectories.
- **Incremental Backup**: Captures only the changes made since the last full backup, including new, modified, or deleted files and directories.

### Requirements:
- **Files** are considered changed if their modification date or size has changed.
- **Directories** are considered changed if their modification date differs.
- The date and time for backup directories are automatically filled with the current date and time in the format `YYYY-MM-DD_HH-MM-SS`.

### Usage:

#### Backup
```sh
my_backup <full | incremental> <workDir> <backupDir>
- **full**: Creates a full backup of the `/work` directory.
- **incremental**: Creates an incremental backup since the last full backup.


# Backup Restore Utility

This utility restores the backup from `<backupDir>` to `<restoreDir>`.

## Parameters:

- **backupDir**: The path to the backup directory where the backup to restore is stored.
- **restoreDir**: The path to the directory where the backup will be restored.

## Example Usage:

### Restore a backup:
```bash
my_restore /backup/2024-01-01_00-00-00 /work

# Error Handling

The utility provides error messages to guide the user when issues occur:

## Missing Permissions
If the program encounters permission issues (e.g., reading/writing files), it will display a message explaining the problem and suggest steps to resolve it.

## Disk Space
If there is not enough space on the disk for the backup, an error message will be shown with suggestions for freeing up space.

## Missing Full Backup for Incremental
If an incremental backup is attempted without a previous full backup, an error will be raised, suggesting to create a full backup first.

## Invalid Paths
If the specified directories for backup or restore do not exist, a clear error message will guide the user to the correct paths.


# Backup and Restore Test Scenarios

## 1. Test Full Backup Creation
**Purpose**: Ensure a full backup is created successfully.

**Steps**:
- Create a full backup from the `/work` directory to the `/backup` directory.
- Verify that all files and subdirectories from `/work` are included in the backup.

---

## 2. Test Restore from Full Backup
**Purpose**: Ensure data is correctly restored from a full backup.

**Steps**:
- Create a full backup of the `/work` directory.
- Restore the backup to a new directory (e.g., `/restore`).
- Verify that the contents of `/restore` match the original `/work` directory.

---

## 3. Test Invalid Paths and Permissions
**Purpose**: Ensure the utility handles invalid paths or insufficient permissions correctly.

**Steps**:
- Try to create a backup or restore from a directory with invalid paths or insufficient permissions.
- Verify that an appropriate error message is shown.

---

## 4. Test Insufficient Disk Space
**Purpose**: Ensure the utility handles low disk space during backup or restore.

**Steps**:
- Simulate low disk space.
- Attempt to create or restore a backup and verify that the utility reports the disk space issue.



