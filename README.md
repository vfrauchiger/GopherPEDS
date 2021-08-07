![alt text](./VinzF.png)

# GopherPEDS
## A GUI Tool for USPTO PEDS 

GopherPEDS is a tool to access the PEDS System of the US Patent and Trademark Office. 

It currently allows for catching the term extension days for an application id, 
early publication number or patent number. When catching the days it also checks for 
terminal disclaimer and gives notice about such disclaimer.

It further allows for downloading of full USPTO File Wrappers. Each file of the File 
Wrapper will be downloaded as separate file in the following format:

            14567345_2015-07-15T00_00_00.000Z_Claims.pdf

            ApplIdFiling_Date&Time_TypeOfDocument.pdf

The files are saved to the location from where you started the application.

To build the application:

_Prerequisites_:
- You need to have Go 1.16+ (golang.org) installed
- You need to have Fyne (Fyne.io) with cmd installed:

        go get fyne.io/fyne/v2
        go get fyne.io/fyne/v2/cmd/fyne


_Steps_:
1. Clone the repo to your preferred location
2. change into the GopherPEDS directory
3. "go build" the application (may be omitted as Fyne does it as well)
4. "fyne package -icon icon.png" to package the application (tested on MacOS Big Sur and Win 10 (20H2))

This app has no intention to have any commercial aspect! Use it or change it. 

*Credits to all GO developers! Credits to the Fyne Team!*

![alt text](./gopherli.png) ![alt text](./fyne.png)

... and may your appropriate God bless the **Cheese Makers**!



