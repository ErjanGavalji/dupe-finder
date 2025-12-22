# Image Backup Consolidator

A tool to identify duplicate images across multiple backup locations and
consolidate them into a single source of truth.

## Purpose

Photos accumulated over years are scattered across multiple backup locations
(external hard disks, CDs, flash drives). With tens of thousands of image files—
including resized versions and various attempts to organize, it gets difficult
to identify what's truly unique, not to mention the disk space wasted.

## The Goal

Analyze files and folders in specified root directories and present a way to
inspect the duplicates. Then, suggest ways to consolidate all photos into a
single source of truth.

## (Current) Idea

The tool will:

- Build a list of all directories (identified by path)
- For each directory, catalog all files
- Each file has a list of duplicates
- Each directory has a list of duplicates above a certain percentage of
  duplication (specified through the command line)

**Scope:** While the primary focus is image files, the tool should process all
files. This helps identify directories that are complete duplicates—if all
images match, a non-image file present in only one location can be flagged for
review.

**Ignore patterns:** Support for excluding system files (thumbs.db, .DS_Store,
etc.) will be included.

## User Experience

The presentation will be user-centric: the entire duplicate structure should be
understandable at a glance, without requiring navigation between pages or
remembering information from different views.
