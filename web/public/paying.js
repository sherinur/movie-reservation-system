document.getElementById('phone').addEventListener('input', function (e) {
    let x = e.target.value.replace(/\D/g, '').match(/(\d{1})(\d{0,3})(\d{0,3})(\d{0,2})(\d{0,2})/);
    e.target.value = '+7' + (x[2] ? ` (${x[2]}` : '') + 
                        (x[3] ? `) ${x[3]}` : '') + 
                        (x[4] ? `-${x[4]}` : '') + 
                        (x[5] ? `-${x[5]}` : '');
});
