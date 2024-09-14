package main

// sumaregi scopes
const (
	// product
	ProductRead  = "pos.product:read"
	ProductWrite = "pos.product:write"
	// customer
	CustomerRead  = "pos.customer:read"
	CustomerWrite = "pos.customer:write"
	// stock
	StockRead   = "pos.stock:read"
	StockWrite  = "pos.stock:write"
	StockChange = "pos.stock-changes:read"
	// transaction
	TransactionRead  = "pos.transaction:read"
	TransactionWrite = "pos.transaction:write"
	// suppliers
	SuppliersRead  = "pos.suppliers:read"
	SuppliersWrite = "pos.suppliers:write"
	// stores
	StoresRead  = "pos.stores:read"
	StoresWrite = "pos.stores:write"
	// staffs
	StaffsRead  = "pos.staffs:read"
	StaffsWrite = "pos.staffs:write"
	// losses
	LossesRead  = "pos.losses:read"
	LossesWrite = "pos.losses:write"
	// orders
	OrdersRead  = "pos.orders:read"
	OrdersWrite = "pos.orders:write"
	// transfers
	TransfersRead  = "pos.transfers:read"
	TransfersWrite = "pos.transfers:write"
	// stock-takings
	StockTakingsRead = "pos.stock-takings:read"
	// order-shipment-stores
	OrderShipmentStoresRead  = "order-shipment.stores:read"
	OrderShipmentStoresWrite = "order-shipment.stores:write"
	// order-shipment-orders
	OrderShipmentOrdersRead  = "order-shipment.orders:read"
	OrderShipmentOrdersWrite = "order-shipment.orders:write"
	// order-shipment-stocks
	OrderShipmentStocksRead = "order-shipment.stocks:read"
)

// sumaregi waiter scopes
const (
	// stores
	WaiterStoresRead  = "waiter.stores:read"
	WaiterStoresPrint = "waiter.stores:print"
	// menues
	WaiterMenuesRead = "waiter.menues:read"
	// orders
	WaiterOrdersRead    = "waiter.orders:read"
	WaiterOrdersWrite   = "waiter.orders:write"
	WaiterOrdersHistory = "waiter.orders:history"
	// reservations
	WaiterReservationsRead  = "waiter.reservations:read"
	WaiterReservationsWrite = "waiter.reservations:write"
	// staff-calls
	WaiterStaffCallsRead = "waiter.staff-calls:read"
)

// sumaregi timecard scopes
const (
	// attendances
	TimecardAttendancesWrite = "timecard.attendances:write"
	// shifts
	TimecardShiftsRead  = "timecard.shifts:read"
	TimecardShiftsWrite = "timecard.shifts:write"
	// salaries
	TimecardSalariesRead  = "timecard.salaries:read"
	TimecardSalariesWrite = "timecard.salaries:write"
	// daily-reports
	TimecardDailyReportsRead  = "timecard.daily-reports:read"
	TimecardDailyReportsWrite = "timecard.daily-reports:write"
	// holidays
	TimecardHolidaysRead  = "timecard.holidays:read"
	TimecardHolidaysWrite = "timecard.holidays:write"
	// stores
	TimecardStoresRead = "timecard.stores:read"
	// staffs
	TimecardStaffsRead = "timecard.staffs:read"
	// settings
	TimecardSettingsRead = "timecard.settings:read"
)
