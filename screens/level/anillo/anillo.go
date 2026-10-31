components {
  id: "point_trigger"
  component: "/screens/level/anillo/anillo.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 30.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite_frente"
  type: "sprite"
  data: "default_animation: \"anillo-frente\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/elements/assets.atlas\"\n"
  "}\n"
  ""
  position {
    z: 2.0
  }
  scale {
    x: 2.0
    y: 2.0
  }
}
embedded_components {
  id: "sprite_fondo"
  type: "sprite"
  data: "default_animation: \"anillo-fondo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/elements/assets.atlas\"\n"
  "}\n"
  ""
  position {
    z: -2.0
  }
  scale {
    x: 2.0
    y: 2.0
  }
}
embedded_components {
  id: "sound1"
  type: "sound"
  data: "sound: \"/assets/sfx/impactMetal_light_000.ogg\"\n"
  ""
}
embedded_components {
  id: "sound2"
  type: "sound"
  data: "sound: \"/assets/sfx/impactMetal_light_001.ogg\"\n"
  ""
}
embedded_components {
  id: "sound4"
  type: "sound"
  data: "sound: \"/assets/sfx/impactMetal_light_003.ogg\"\n"
  ""
}
embedded_components {
  id: "sound3"
  type: "sound"
  data: "sound: \"/assets/sfx/impactMetal_light_002.ogg\"\n"
  ""
}
embedded_components {
  id: "sound5"
  type: "sound"
  data: "sound: \"/assets/sfx/impactMetal_light_004.ogg\"\n"
  ""
}
