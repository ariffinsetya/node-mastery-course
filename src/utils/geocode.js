const request = require('postman-request')
const geocode = (searchText, callback) => {
    const url = `https://api.mapbox.com/geocoding/v5/mapbox.places/${encodeURIComponent(searchText)}.json?access_token=pk.eyJ1IjoiYXJpZmZpbnNldHlhIiwiYSI6ImNrcmxpdnBzNzEwYnIyb2xpNHVkb3FocTkifQ.susYtIKrcpTPBfo27Y6ikA`
    request(
        {url, json: true},
        (err,{body}) => {
            if(err){
                callback('Unable to connect!', undefined)
            } else if(body.features.length === 0){
                callback('No valid data found', undefined)
            } else {
                const data = {
                    longitude: body.features[0].center[0],
                    latitude: body.features[0].center[1],
                    location: body.features[0].place_name
                }
                callback(undefined,data)
            }
        }
    )
}

module.exports = geocode