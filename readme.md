# Todaytvseries organizers

The main purpose of this project is to Extracts and Organizes movies downloaded from [todaytvseries.com](todaytvseries.com). It can also be used to 
extract and organize different file types based on the configuration settings

This project was written in Go. Just a Noob learning new things :).

## Usage

            $ go get github.com/kayslay/todaytvseries_organizer
            $ cd $GOPATH/src/github.com/kayslay/todaytvseries_organizer
            $ make
            $ cp config.json /path/of/downloaded/rar/files
            $ cd /path/of/downloaded/rar/files
            $ tv_org

This project is learning process for me.

## Config File Settings

- **deleteAfter**: if true deletes the `RAR` file after extraction is complete.

- **moveDir**: the root directory to place the file.

- **path**: the path where the `RAR` files are located.

- **ext**: the extension to match. defaults to rar

- **ext**: the extension of the file to get out of the compressed file

- **folderName**: the regular expression used to generate the file folder form the file name. It generates the folder based on the grouping extracted of the regexp.

- **workerCount**: the amount of concurrent workers to run at the same time. workerCount must be greater than 0. It defaults to 1.