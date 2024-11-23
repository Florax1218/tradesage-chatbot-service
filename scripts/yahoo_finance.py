# scripts/yahoo_finance.py

import yfinance as yf
import sys
import json
from datetime import datetime, timedelta

def get_stock_data(symbol):
    try:
        ticker = yf.Ticker(symbol)
        # Obtain real-time data
        info = ticker.info
        # Retrieve historical data (last 7 days)
        end_date = datetime.now()
        start_date = end_date - timedelta(days=7)
        hist = ticker.history(start=start_date, end=end_date)
        
        # Build return data
        result = {
            "current_data": {
                "symbol": symbol,
                "price": info.get("regularMarketPrice", 0),
                "volume": info.get("regularMarketVolume", 0),
                "timestamp": int(datetime.now().timestamp())
            },
            "historical_data": []
        }
        
        # Add historical data
        for index, row in hist.iterrows():
            result["historical_data"].append({
                "date": index.strftime("%Y-%m-%d"),
                "price": row["Close"],
                "volume": row["Volume"],
                "timestamp": int(index.timestamp())
            })
        
        print(json.dumps(result))
        return 0
    except Exception as e:
        print(json.dumps({"error": str(e)}), file=sys.stderr)
        return 1

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({"error": "Please provide a stock symbol"}), file=sys.stderr)
        sys.exit(1)
        
    sys.exit(get_stock_data(sys.argv[1]))