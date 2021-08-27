# autonotes

Automation tool to autocapture screenshots and join them with a supplied .srt or .txt file and output a notes file in markdown.

## Usage

For the sake of simplicity, `mage` was used.

Show list of options:
```bash
mage
```

outputs:
```bash
❯ mage
Mage is a make-like command runner.  See https://magefile.org for full docs.

Targets:
  notes:generate    .md notes file - usage: mage generate <file_and_folder_name_must_have_the_same_name>
  screen:capture    screen - usage: mage screen:capture <start_delay_seconds> <screen_index> <capture_interval_ms> <jpeg_or_png> <output_folder_name>
  screen:test       screen capture all screens and saves them to a "screen_test" folder with the number of the screen
```

### screen test
Tests screen capture on all screens and outputs the images to "output/screen_test".

Use this to find the `index` of the screen you want to capture in case you have multiple monitors.
```bash
mage screen:test
```

### screen capture
Captures screenshots from a monitor, within an interval (if the image is different - overall works well but your usage may vary), and saves them to a folder with a timestamp on the name.

```bash
mage screen:capture <start_delay_seconds> <screen_index> <capture_interval_ms> <jpeg_or_png> <output_folder_name>
```

example usage:
```bash
mage screen:capture 0 0 5000 jpeg my-presentation-notes
```

### generate notes
This assumes that a folder with screenshots was already created (refer to **screen capture** above).

**IMPORTANT:** you have to put the input file in a folder named `input` on the root of this project.

Both the folder and the `.srt` file should have the same name, example folder: `my-presentation-notes` and example .srt file: `my-presentation-notes.srt`.
```bash
mage generate <file_and_folder_name_must_have_the_same_name>
```

example usage:
```bash
mage generate my-presentation-notes
```

## TODO

- [ ] support txt files
- [ ] integrate with google or aws speech-to-text solutions
- [ ] output with obsidian image links format
- [ ] test on Windows
- [ ] add name of file as title in md file
