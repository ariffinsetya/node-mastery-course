const request = require('postman-request')
const forecast = (longitude, latitude, callback) => {
    const url = `http://api.weatherstack.com/current?access_key=a9da026598a7d33974998ccdb9c1736d&query=${latitude},${longitude}`
    request({ url, json: true },
        (err, { body }) => {
            if (err) {
                callback('Unable to connect!', undefined)
            } else if (body.error !== undefined) {
                callback(body.info, undefined)
            } else {
                body.current.forecastData = "Weather is " + body.current.weather_descriptions[0] + " Temperature is " + body.current.temperature + " degrees, but it feels like " + body.current.feelslike + " degrees now"
                callback(undefined, body.current)
            }
        }
    )
}

module.exports = forecast