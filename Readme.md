# Image Backup Consolidator

A tool to identify duplicate images across multiple backup locations and
consolidate them into a single source of truth.

## Purpose

I have a lot of photos gathered through the years. I have made tons of backups
to all kinds of external hard disks, CDs or flash drives.

I want to have those consolidated in a single location, which will be the
"source of truth" and use a decent backup software (probably Restic). Then, free
all the other hardware from those photos for other uses.

With tens of thousands of image files, some resized versions of the originals, I
need to build a map the duplicates, analyze the trees of directories (as some
images are inside different subdirectories as earlier attempts to organize all)
and finally, present a suggestion for the easiest action. Applying the action is
an option as well.

## (Current) Idea

I will not try to reconstruct the file system for the time being.

I will have a list of all the directories (identified by a path). Each directory
will have a list of files. Right now, this is for image files only, but I'd
rather read any file, as that might lead to the deletion of a non-image file
that exists in one folder, but doesn't exist in another, simply because all the
image files were the same. Of course, it would be good to specify ignored
filename patterns (thumbs.db, .DS_Store, etc)
