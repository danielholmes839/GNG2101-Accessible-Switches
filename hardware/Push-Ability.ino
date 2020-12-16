#define pushButton1 2
#define pushButton2 3
#define pushButton3 4
#define pushButton4 5
#define pushButton5 6

bool pressed1 = false, pressed2 = false, pressed3 = false, pressed4 = false, pressed5 = false;

void setup() {
  // put your setup code here, to run once:
  pinMode(pushButton1, INPUT_PULLUP); //set pin 2 as button with pullup resistor
  pinMode(pushButton2, INPUT_PULLUP); //set pin 3 as button with pullup resistor
  pinMode(pushButton3, INPUT_PULLUP); //set pin 4 as button with pullup resistor
  pinMode(pushButton4, INPUT_PULLUP); //set pin 5 as button with pullup resistor
  pinMode(pushButton5, INPUT_PULLUP); //set pin 6 as button with pullup resistor

  Serial.begin(9600); //start serial comms 9600 baud

}

void loop() {
  // put your main code here, to run repeatedly:

  if (!digitalRead(pushButton1) && !pressed1){
    Serial.write(1);
    pressed1 = true;
  }
    else if (digitalRead(pushButton1))
    {
      pressed1 = false;
    }
  
  if (!digitalRead(pushButton2) && !pressed2){
    Serial.write(2);
    pressed2 = true;
  }
    else if (digitalRead(pushButton2))
    {
      pressed2 = false;
    }

  if (!digitalRead(pushButton3) && !pressed3){
  Serial.write(3);
  pressed3 = true;
  }
    else if (digitalRead(pushButton3))
    {
      pressed3 = false;
    }

  if (!digitalRead(pushButton4) && !pressed4){
  Serial.write(4);
  pressed4 = true;
  }
    else if (digitalRead(pushButton4))
    {
      pressed4 = false;
    }

  if (!digitalRead(pushButton5) && !pressed5){
  Serial.write(5);
  pressed5 = true;
  }
    else if (digitalRead(pushButton5))
    {
      pressed5 = false;
    }
  
  delay(1);

}
