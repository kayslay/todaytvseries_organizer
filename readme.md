# Todaytvseries organizers

Extracts  and Organizes movies downloaded from [todaytvseries.com](todaytvseries.com).

Written in Go. 

## Usage 

            $ go get github.com/kayslay/todaytvseries_organizer
            $ cd $GOPATH/src/github.com/kayslay/todaytvseries_organizer
            $ cp config.json /path/of/downloaded/rar/files
            $ cd /path/of/downloaded/rar/files
            $ tv_org


This project is  learning process for me.

## Todo

- add concurrent extractions to the project.

## Config File Settings

- **deleteAfter**: if true deletes the `RAR` file after extarction is complete.

- **moveDir**: the root directory to place the file.

- **path**: the path where the `RAR` files are located.

- **ext**: the extension to match. defaults to rar

- **ext**: the extension of the file to get out of the compressed file

- **folderName**: the regular expression used to generate the file folder form the file name. It generates the folder based on the grouping extracted of the regexp.

