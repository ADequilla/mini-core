package database

var QueryGetAllFees = "SELECT * FROM mfs.m_fee_Structure"
var QueryGetSpecifiedFee = "SELECT * FROM mfs.m_fee_Structure WHERE fee_id=?"
var QueryCreateToken = "INSERT INTO mfs.app_token (appid,username,firstname,lastname,token,updated_at_unix) VALUES(?,?,?,?,?,?)"
var QueryUpdateToken = "UPDATE mfs.app_token SET appid=?,username=?,firstname=?,lastname=?,token=? ,updated_at_unix=? WHERE appid=? AND username=? "
var QueryUpdateGetExp = "UPDATE mfs.app_token SET updated_at_unix=? WHERE  appId=? AND username=? "
