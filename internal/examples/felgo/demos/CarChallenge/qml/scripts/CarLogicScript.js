var hits = 0;
// Time in s how long after firing a rocket or a mine the car is invincible. This time depends on the speed of the rocket and the car (set it so the rocket will definetely be away from the car)!
var DELAY_CAR_VULNERABLE_TO_JUSTSHOT_WEAPON = 3;

// this should be private - is there a way not to allow access to this variable?
// set it to -delay, so at first press immediately a rocket can be created even if core time is 0
var firetime=-delayMultipleFireActions;

// needed to avoid being invulnerable to own fired weapons as well
var justFiredWeaponId;

// this may be set to rocket_green or mine after a collision with a weaponbox
// for testing a specific type of weapon, this could be set to any and the car will have it at program start!
var activeWeapon = "";

function onInitialize() {
    //hits = 0
}
