const myText = document.getElementById("myText");
const mySubmit = document.getElementById("mySubmit");
const resultElement = document.getElementById("resultElement");
let age;

mySubmit.onclick = function(){
    age = myText.value;
    age = Number(age)

    if(age >= 100){
        
        resultElement.textContent = `Your are too old to enter this site`;
        
    }
    else if(age == 0){
        
                resultElement.textContent = `you cant enter . you are just born`;

    }
    else if(age >= 18){
        
        resultElement.textContent = `You are old enough to enter this site`;

        
    }
    else if(age < 0){
        
         resultElement.textContent = `your age cant be below 0`;
        
    }
    else{
        console.log("");
         resultElement.textContent = `you must be 18+ to enter this site`;
        
    }

    }