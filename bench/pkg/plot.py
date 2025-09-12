import sys
import json
import matplotlib.pyplot as plt
import matplotlib.ticker as mticker

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

    if not x or not datas or len(datas) == 0:
        print("Error: JSON must contain 'x' and a non-empty 'datas' list.")
        sys.exit(1)

    # Expect only one data series
    series = datas[0]
    series_name = series.get("Name", "Unnamed Series")
    series_data = series.get("Data")

    if not series_data:
        print("Error: The data series must contain a 'Data' list.")
        sys.exit(1)

    fig, ax = plt.subplots(figsize=(12, 8))

    ax.plot(x, series_data, label=series_name)

    ax.set_title(f'Benchmark: {series_name}')
    ax.set_xlabel('Run Number')
    ax.set_ylabel(series_name)
    ax.grid(True)

    # Set a reasonable number of x-axis ticks
    ax.xaxis.set_major_locator(mticker.MaxNLocator(nbins=10, integer=True))
    fig.autofmt_xdate(rotation=45) # Auto-format x-axis labels

    plt.tight_layout()

    plt.savefig(output_filename, dpi=150)
    # print(f"Chart saved to {output_filename}") # Comment out to not pollute stdout for Go

if __name__ == "__main__":
    main()
