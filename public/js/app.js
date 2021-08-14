console.log('Client side javascript file is loaded!')

const searchForm = document.querySelector('form')

const searchEl = document.querySelector('input')

const msgOneEl = document.querySelector('#msg-1')
const msgTwoEl = document.querySelector('#msg-2')

searchForm.addEventListener('submit', (e) => {
    e.preventDefault();

    msgOneEl.textContent = 'Loading...'
    msgTwoEl.textContent = ''
    fetch(`/weather?address=${searchEl.value}`).then((res) => {
        res.json().then((data) => {
            if (data.error) {
                msgOneEl.textContent = "Error"
                msgTwoEl.textContent = data.error
            } else {
                msgOneEl.textContent = data.location
                msgTwoEl.textContent = data.forecast
            }
        })
    })
})