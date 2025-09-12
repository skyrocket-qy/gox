import sys
import json
import matplotlib.pyplot as plt

def main():
    if len(sys.argv) < 2:
        print("Usage: python plot.py <output_filename>")
        sys.exit(1)

    output_filename = sys.argv[1]

    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError:
        print("Error: Invalid JSON received from stdin.")
        sys.exit(1)

    x = data.get("x")
    datas = data.get("datas")

    if not x or not datas:
        print("Error: JSON must contain 'x' and 'datas' keys.")
        sys.exit(1)

    fig, ax = plt.subplots(figsize=(12, 8))

    for series in datas:
        series_name = series.get("Name")
        series_data = series.get("Data")
        if series_name and series_data:
            ax.plot(x, series_data, label=series_name)

    ax.set_title('Benchmark Results')
    ax.set_xlabel('Run Number')
    ax.set_ylabel('Value')
    ax.legend()
    ax.grid(True)

    # Rotate x-axis labels to prevent overlap if there are many runs
    plt.xticks(rotation=45, ha='right')
    plt.tight_layout()

    plt.savefig(output_filename, dpi=150)
    print(f"Chart saved to {output_filename}")

if __name__ == "__main__":
    main()
