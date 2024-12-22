import json
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
from websocket import create_connection, WebSocketApp
import threading

# Data storage for plotting
trajectories = {}

# Lock for thread-safe data access
data_lock = threading.Lock()

# WebSocket message handler
def on_message(ws, message):
    print(f"Received: {message}")
    global trajectories
    # Parse incoming message
    point = json.loads(message)
    _id = point["id"]
    with data_lock:
        if _id not in trajectories:
            trajectories[_id] = {"x": [], "y": []}
        trajectories[_id]["x"].append(point["x"])
        trajectories[_id]["y"].append(point["y"])

# WebSocket error handler
def on_error(ws, error):
    print(f"WebSocket Error: {error}")

# WebSocket close handler
def on_close(ws, close_status_code, close_msg):
    print("WebSocket Closed")

# WebSocket open handler
def on_open(ws):
    print("WebSocket Connection Opened")

# WebSocket setup
websocket_url = "ws://127.0.0.1:8090/ws"  # Replace with your WebSocket server URL
ws = WebSocketApp(websocket_url,
                  on_message=on_message,
                  on_error=on_error,
                  on_close=on_close,
                  on_open=on_open)

# Run WebSocket in a separate thread
def run_websocket():
    ws.run_forever()

thread = threading.Thread(target=run_websocket)
thread.daemon = True
thread.start()

# Set up the plot
fig, ax = plt.subplots(figsize=(10, 10))
lines = {}

ax.set_xlim(0, 100)  # Adjust based on your data range
ax.set_ylim(0, 100)  # Adjust based on your data range
ax.set_title("Live 2D Grid Map with Trajectories")
ax.set_xlabel("X Coordinate")
ax.set_ylabel("Y Coordinate")
ax.grid(True)

# Update function for animation
def update(frame):
    global trajectories
    with data_lock:
        for _id, trajectory in trajectories.items():
            if _id not in lines:
                # Create a new line for each ID
                lines[_id] = ax.plot([], [], marker="o", label=f"ID {_id}")[0]
                ax.legend()
            line = lines[_id]
            line.set_data(trajectory["x"], trajectory["y"])
    return lines.values()

# Create the animation
ani = FuncAnimation(fig, update, interval=500, blit=True)

plt.show()
