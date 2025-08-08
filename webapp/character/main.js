let currentCharacter = null;

async function loadCharacter() {
  const id = document.getElementById("charId").value.trim();
  if (!id) {
    alert("Please enter a character ID.");
    return;
  }

  try {
    const res = await fetch(`${SERVER_URL}/char/?id=${id}`);
    if (!res.ok) {
      throw new Error(`Server returned ${res.status}`);
    }
    currentCharacter = await res.json();
    populateForm(currentCharacter);
  } catch (err) {
    console.error("Error loading character:", err);
    alert(`${SERVER_URL}/char/?id=${id}`)
    //alert("Failed to load character with id Check the ID and try again.");
  }
}

function populateForm(data) {
  try {
    document.getElementById("name").value = data.name || "";
    document.getElementById("alignment").value = data.alignment || "";
    document.getElementById("hp").value = data.status.hp || 0;
    document.getElementById("max_hp").value = data.status.max_hp || 0;
    document.getElementById("armor_class").value = data.status.armor_class || 0;
    document.getElementById("speed").value = data.status.speed || 0;
    document.getElementById("initiative").value = data.status.initiative || 0;
    document.getElementById("strength").value = data.attributes.strength || 0;
    document.getElementById("dexterity").value = data.attributes.dexterity || 0;
    document.getElementById("constitution").value = data.attributes.constitution || 0;
    document.getElementById("intelligence").value = data.attributes.intelligence || 0;
    document.getElementById("wisdom").value = data.attributes.wisdom || 0;
    document.getElementById("charisma").value = data.attributes.charisma || 0;
  } catch (err) {
    console.error("Error populating form:", err);
    alert("Character data is invalid or incomplete.");
  }
}

async function saveCharacter() {
  if (!currentCharacter) {
    alert("No character loaded.");
    return;
  }

  try {
    currentCharacter.name = document.getElementById("name").value;
    currentCharacter.alignment = document.getElementById("alignment").value;
    currentCharacter.status.hp = parseInt(document.getElementById("hp").value);
    currentCharacter.status.max_hp = parseInt(document.getElementById("max_hp").value);
    currentCharacter.status.armor_class = parseInt(document.getElementById("armor_class").value);
    currentCharacter.status.speed = parseInt(document.getElementById("speed").value);
    currentCharacter.status.initiative = parseInt(document.getElementById("initiative").value);
    currentCharacter.attributes.strength = parseInt(document.getElementById("strength").value);
    currentCharacter.attributes.dexterity = parseInt(document.getElementById("dexterity").value);
    currentCharacter.attributes.constitution = parseInt(document.getElementById("constitution").value);
    currentCharacter.attributes.intelligence = parseInt(document.getElementById("intelligence").value);
    currentCharacter.attributes.wisdom = parseInt(document.getElementById("wisdom").value);
    currentCharacter.attributes.charisma = parseInt(document.getElementById("charisma").value);

    const res = await fetch(`${SERVER_URL}/char/?id=${currentCharacter.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(currentCharacter)
    });

    if (!res.ok) {
      throw new Error(`Server returned ${res.status}`);
    }

    alert("Character updated successfully!");
  } catch (err) {
    console.error("Error saving character:", err);
    alert("Failed to save character. Please try again.");
  }
}