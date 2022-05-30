import crypto from "crypto";
import { faker } from "@faker-js/faker";

const users = Array.from({ length: 10000 }).map((_, i) => ({
  id: i,
  firstName: faker.name.firstName(),
  lastName: faker.name.lastName(),
  email: faker.internet.email(),
  phone: faker.phone.phoneNumber(),
  address: faker.address.streetAddress(),
  city: faker.address.city(),
  state: faker.address.state(),
  zip: faker.address.zipCode(),
  company: faker.company.companyName(),
  avatar: faker.internet.avatar(),
}));

const flagIndex = crypto.randomInt(0, 10000);
users[flagIndex].company = "HZU18{9b36fa+NICE_LIST_0e7635b84cefe035}";

console.log(JSON.stringify(users));
