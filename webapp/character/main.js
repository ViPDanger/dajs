let characters = [];
let accessToken = localStorage.getItem("accessToken") || "";
let refreshToken = localStorage.getItem("refreshToken") || "";

async function loadCharacters() {
  if (window.__TOKENS__) {
  localStorage.setItem("accessToken", window.__TOKENS__.access_token);
  localStorage.setItem("refreshToken", window.__TOKENS__.refresh_token);
  delete window.__TOKENS__;
}
  cancelEdit()
  try {
    const res = await apiFetch(`${SERVER_URL}${REQUEST}`);
    if (!res.ok) throw new Error("Ошибка загрузки персонажей");
    if (res.status != 204){
    characters = await res.json(); // массив
    }
    console.log("characters: ",characters)
    renderCharacters();
  } catch (err) {
    alert(`❌ ${err.message}`);
  }
}

function renderCharacters() {
  cancelEdit()
  const container = document.getElementById("characterList");
  console.log("container: ",container)
  container.innerHTML = "";
  characters.forEach((char, index) => {
    const div = document.createElement("div");
    div.className = "char-card";
    div.innerHTML = `
      <button class="btn" onclick="editCharacter(${index})">${char.name}</button>
    `;
    container.appendChild(div);
  });
}

async function addCharacter() {
  const name = prompt("Имя нового персонажа:");
  if (!name) return;
  let newChar = {
    id: '',
    name,
    status: { hp: 6, max_hp: 6, temporary_hp: 0, armor_class: 10, speed: 30, initiative: 0 },
    attributes: {strength: 10, dexterity: 10, constitution: 10, intelligence: 10, wisdom: 10, charisma: 10}, // можно заполнить по умолчанию
    alignment: "",
    abilities: [],
    spells: [],
    tags: []
  };

  try {
    const res = await apiFetch(`${SERVER_URL}${REQUEST}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newChar)
    });
    if (!res.ok) throw new Error("Ошибка добавления персонажа");
    const text = await res.text();
    console.log("post resp: ",res.status,text);
    const body = JSON.parse(text);
    console.log("parse result: ",body);
    newChar.id = body.id;
    console.log("newChar ",newChar)
    characters.push(newChar);
    renderCharacters();
  } catch (err) {
    alert(`❌ ${err.message}`);
  }
}
function editCharacter(index) {
  cancelEdit()
  const char = characters[index];
  document.getElementById("editId").value = char.id;
  document.getElementById("editName").value = char.name;
  document.getElementById("editHp").value = char.status.hp;
  document.getElementById("editMaxHp").value = char.status.max_hp;
  document.getElementById("editAC").value = char.status.armor_class;
  document.getElementById("editSpeed").value = char.status.speed;
  document.getElementById("editInitiative").value = char.status.initiative;

  document.getElementById("attrStrength").value = char.attributes.strength;
  document.getElementById("attrDexterity").value = char.attributes.dexterity;
  document.getElementById("attrConstitution").value = char.attributes.constitution;
  document.getElementById("attrIntelligence").value = char.attributes.intelligence;
  document.getElementById("attrWisdom").value = char.attributes.wisdom;
  document.getElementById("attrCharisma").value = char.attributes.charisma;
  document.getElementById("editor").style.display = "block";
}

function cancelEdit() {
  console.log("editor was canceled")
  document.getElementById("editor").style.display = "none";
}


async function saveCharacter(event) {
  event.preventDefault();

  const id = document.getElementById("editId").value;
  const charIndex = characters.findIndex(c => c.id === id);
  if (charIndex === -1) return;

  const updatedChar = {
    ...characters[charIndex],
    name: document.getElementById("editName").value,
    status: {
      ...characters[charIndex].status,
      hp: parseInt(document.getElementById("editHp").value),
      max_hp: parseInt(document.getElementById("editMaxHp").value),
      armor_class: parseInt(document.getElementById("editAC").value),
      speed: parseInt(document.getElementById("editSpeed").value),
      initiative: parseInt(document.getElementById("editInitiative").value)
    },
    attributes: {
      strength: parseInt(document.getElementById("attrStrength").value),
      dexterity: parseInt(document.getElementById("attrDexterity").value),
      constitution: parseInt(document.getElementById("attrConstitution").value),
      intelligence: parseInt(document.getElementById("attrIntelligence").value),
      wisdom: parseInt(document.getElementById("attrWisdom").value),
      charisma: parseInt(document.getElementById("attrCharisma").value)
    }
  };

  console.log("saveCharacter updated character: ", updatedChar)
  try {
    const res = await apiFetch(`${SERVER_URL}${REQUEST}?id=${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(updatedChar)
    });
    if (!res.ok) throw new Error("Ошибка обновления персонажа");
    characters[charIndex] = updatedChar;
    console.log("PUT request: ",res.status,await res.text())
    renderCharacters();
  } catch (err) {
    alert(`❌ ${err.message}`);
  }``
}

async function deleteCharacter() {
  console.log("DELETE")
  const id = document.getElementById("editId").value;
  const index = characters.findIndex(c => c.id === id);
  console.log("delete index ",index)
  if (index === -1) return;
  const char = characters[index];
  if (!confirm(`Удалить персонажа ${char.name}?`)) return;
  try {
    const res = await apiFetch(`${SERVER_URL}${REQUEST}?id=${char.id}`, { method: "DELETE" });
    if (!res.ok) throw new Error("Ошибка удаления персонажа");
    console.log("delete resp ",res.status,await res.text())
    console.log("splice: ",characters.splice(index,1));
    renderCharacters();
    
  } catch (err) {
    alert(`❌ ${err.message}`);
  }
}


async function apiFetch(url, options = {}) {
  options.headers = {
    ...(options.headers || {}),
    "Content-Type": "application/json",
    "Authorization": `Bearer ${accessToken}`
  };

  let res = await fetch(url, options);

  // Если токен истёк → пытаемся обновить
  if (res.status === 401 && refreshToken) {
    console.warn("Access token expired, refreshing...");
    const refreshed = await refreshAccessToken();
    if (!refreshed) throw new Error("Не удалось обновить токен");
    options.headers["Authorization"] = `Bearer ${accessToken}`;
    res = await fetch(url, options); // повторный запрос
  }

  return res;
}


async function refreshAccessToken() {
  try {
    const res = await fetch(`${SERVER_URL}/auth/refresh`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ refresh_token: refreshToken })
    });
    if (!res.ok) return false;
    const data = await res.json();
    accessToken = data.Access.access_token;
    refreshToken = data.Refresh.refresh_token || refreshToken; // иногда refresh не меняется
    localStorage.setItem("accessToken", accessToken);
    localStorage.setItem("refreshToken", refreshToken);
    console.log("🔄 Token refreshed");
    return true;
  } catch (err) {
    console.error("Ошибка обновления токена", err);
    return false;
  }
}