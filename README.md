# JPEG Image Checker for broken jpegs

This simple tool will check if it can open and parse the file as JPEG, also if the pixel at the bottom right corner at (xsize-10,ysize-10) coordinates is not a given grey colour. This logic is way too simple but really effective, the success rate is enough for firefighting.

Running it on image that got damaged in automated image compression:

```sh
$ jpgcheck file.jpg
file.jpg Error: broken image (grey bottom)
```

Some other output

```sh
$ jpgcheck broken1.jpg
broken1.jpg Error: invalid JPEG format: missing 0xff00 sequence

$ jpgcheck broken2.jpg
broken2.jpg Error: invalid JPEG format: unknown marker

$ jpgcheck broken3.jpg
broken3.jpg Error: invalid JPEG format: DHT has wrong length

$ jpgcheck broken4.jpg
broken4.jpg Error: invalid JPEG format: missing SOS marker

$ jpgcheck broken5.jpg
broken5.jpg Error: invalid JPEG format: bad Huffman code

$ jpgcheck broken6.jpg
broken6.jpg Error: invalid JPEG format: too many coefficients

$ jpgcheck broken7.jpg
broken7.jpg Error: invalid JPEG format: missing SOF marker
```

No message means good jpeg:

```sh
$ jpgcheck good.jpg
```

Smaller than 10x10 jpgs will also error and some accidental false negatives when the checked pixel is #757575 might happen.

Errors are reported on stderr.
