export interface BazaarItemHistory {
    _id: string;
    hypixel_product_id: string;
    last_updated: number;
    history_24h: BazaarItemHistoryData[];
    history_daily: BazaarItemHistoryData[];
}

export interface BazaarItemHistoryData {
    time: number;
    sell_price: number;
    buy_price: number;
}

export interface BazaarItem {
    _id: string;
    hypixel_product_id: string;
    last_updated: number;
    sell_data: BazaarItemBuySellSummary[];
    buy_data: BazaarItemBuySellSummary[];
    sell_price: number;
    sell_volume: number;
    sell_moving_week: number;
    sell_orders: number;
    buy_price: number;
    buy_volume: number;
    buy_moving_week: number;
    buy_orders: number;
    margin: number;
    margin_percent: number;
}

export interface BazaarItemBuySellSummary {
    amount: number;
    pricePerUnit: number;
    orders: number;
}

export interface BazaarTopItem {
    _id: string;
    hypixel_product_id: string;
    buy_price: number;
    buy_volume: number;
    margin: number;
    margin_percent: number;
    sell_price: number;
    sell_volume: number;
}