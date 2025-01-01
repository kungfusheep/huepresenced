# Hue Presence Detector

This Go program detects presence using a Philips Hue motion sensor and prevents your MacOS machine from sleeping or activating the screensaver by simulating user activity.

## Prerequisites

- A Go compiler
- Philips Hue Bridge and Motion Sensor
- macOS (for `caffeinate` command)

## Installation

Install with go get:

```sh 
go get github.com/kungfusheep/huepresenced 
```

## Usage

1. Set the environment variables for your Philips Hue Bridge IP and username:
    ```sh
    export HUEIP=<your-bridge-ip>
    export HUESERNAME=<your-username>
    ```

2. Run the program:
    ```sh
    huepresenced -sensor <sensor-id> -log
    ```

    - `-sensor`: The ID of the Hue motion sensor to monitor. 
    - `-log`: (Optional) Enable logging of presence detection.

## How It Works

- The program continuously checks the presence state of the specified Hue motion sensor every 5 seconds.
- If presence is detected, it runs the `caffeinate -dimsut 1` command to simulate user activity, preventing the macOS machine from sleeping or activating the screensaver.

## Example

```sh
export HUEIP=192.168.1.2
export HUESERNAME=your-hue-username
./huepresenced -sensor 5 -log
```

This will monitor the Hue motion sensor with ID `3` and log presence detection events.


## How to Get Hue Bridge IP and Username

1. Find the IP address of your Philips Hue Bridge:
    - Open the Hue app on your phone.
    - Go to Settings > Hue Bridges.
    - Tap on the "i" icon next to your bridge.
    - Look for the "IP Address" field.

2. Get the username for your Hue Bridge:
    - Press the button on your Hue Bridge.
    - Within 30 seconds, run `curl -X POST "http://$HUEIP/api" -d '{"devicetype":"huepresenced#macos"}'` in your terminal. This will create a new username. 

## How to Find the Sensor ID

1. Find the sensor ID of your Philips Hue Motion Sensor:
    - Run `curl -X GET "http://$HUEIP/api/$HUESERNAME/sensors"` in your terminal.
    - Look for the sensor with the type `ZLLPresence` or the name you recognize - the ID is the key

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```


